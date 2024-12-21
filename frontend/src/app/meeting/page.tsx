'use client';

import React, { useEffect, useRef, useState } from "react";
import { Video, VideoOff, Mic, MicOff, Monitor, Volume2, VolumeX, LogOut } from 'lucide-react';
import { useRouter } from 'next/navigation';

// Defina um tipo para os participantes
type Participant = {
  id: string;
  name: string;
};

const Meeting: React.FC = () => {
  const router = useRouter();

  const localVideoRef = useRef<HTMLVideoElement>(null);
  const remoteVideoRef = useRef<HTMLVideoElement>(null);
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const peerConnection = useRef<RTCPeerConnection | null>(null);

  const [isVideoEnabled, setIsVideoEnabled] = useState(true);
  const [isAudioEnabled, setIsAudioEnabled] = useState(true);
  const [isTabMuted, setIsTabMuted] = useState(false);
  const [isScreenSharing, setIsScreenSharing] = useState(false);

  const [participants, setParticipants] = useState<Participant[]>([
    { id: 'local', name: 'Você' }
  ]);

  const getGridLayout = (participantCount: number) => {
    if (participantCount <= 4) {
      return `grid-cols-${participantCount}`;
    }
    return 'grid-cols-4';
  };

  useEffect(() => {
    const initWebRTC = async () => {
      try {
        const localStream = await navigator.mediaDevices.getUserMedia({
          video: true,
          audio: true,
        });
        
        if (localVideoRef.current) {
          localVideoRef.current.srcObject = localStream;
        }

        const pc = new RTCPeerConnection({
          iceServers: [
            { urls: "stun:stun.l.google.com:19302" },
            { 
              urls: "turn:localhost:3478",
            }
          ],
        });

        localStream.getTracks().forEach((track) => {
          pc.addTrack(track, localStream);
        });

        pc.ontrack = (event) => {
          if (remoteVideoRef.current) {
            remoteVideoRef.current.srcObject = event.streams[0];
          }
        };

        pc.onicecandidate = (event) => {
          if (event.candidate && socket) {
            socket.send(
              JSON.stringify({
                type: "candidate",
                payload: event.candidate,
              })
            );
          }
        };

        peerConnection.current = pc;
      } catch (error) {
        console.error('Erro ao inicializar WebRTC:', error);
        alert('Erro ao acessar câmera/microfone. Por favor, verifique as permissões.');
      }
    };

    initWebRTC();

    const ws = new WebSocket("ws://localhost:8080/ws?id=user1&room=room1");
    
    ws.onerror = (error) => {
      console.error('Erro na conexão WebSocket:', error);
      alert('Erro na conexão. Tentando reconectar...');
    };

    ws.onclose = () => {
      console.log('Conexão WebSocket fechada');
      setTimeout(() => {
        console.log('Tentando reconectar...');
        setSocket(new WebSocket("ws://localhost:8080/ws?id=user1&room=room1"));
      }, 5000);
    };

    ws.onmessage = async (event) => {
      const message = JSON.parse(event.data);

      if (message.type === "offer") {
        await peerConnection.current?.setRemoteDescription(new RTCSessionDescription(message.payload));
        const answer = await peerConnection.current?.createAnswer();
        if (answer) {
          await peerConnection.current?.setLocalDescription(answer);
          ws.send(
            JSON.stringify({
              type: "answer",
              payload: answer,
            })
          );
        }
      } else if (message.type === "answer") {
        await peerConnection.current?.setRemoteDescription(new RTCSessionDescription(message.payload));
      } else if (message.type === "candidate") {
        await peerConnection.current?.addIceCandidate(new RTCIceCandidate(message.payload));
      }
    };

    setSocket(ws);

    return () => {
      ws.close();
    };
  }, []);

  const toggleVideo = () => {
    const videoTrack = (localVideoRef.current?.srcObject as MediaStream | null)
        ?.getTracks()
        .find(track => track.kind === 'video');
    if (videoTrack) {
        videoTrack.enabled = !isVideoEnabled;
        setIsVideoEnabled(!isVideoEnabled);
    }
  };

  const toggleAudio = () => {
    const audioTrack = (localVideoRef.current?.srcObject as MediaStream | null)
        ?.getTracks()
        .find(track => track.kind === 'audio');
    if (audioTrack) {
        audioTrack.enabled = !isAudioEnabled;
        setIsAudioEnabled(!isAudioEnabled);
    }
  };

  const toggleTabAudio = () => {
    if (remoteVideoRef.current) {
      remoteVideoRef.current.muted = !isTabMuted;
      setIsTabMuted(!isTabMuted);
    }
  };

  const shareScreen = async () => {
    try {
      const screenStream = await navigator.mediaDevices.getDisplayMedia({ video: true });
      const videoTrack = screenStream.getVideoTracks()[0];
      const sender = peerConnection.current
        ?.getSenders()
        .find(s => s.track?.kind === 'video');
      if (sender) {
        sender.replaceTrack(videoTrack);
      }
    } catch (error) {
      console.error('Erro ao compartilhar tela:', error);
    }
  };

  const leaveMeeting = () => {
    if (socket) {
      socket.close();
    }
    if (localVideoRef.current?.srcObject) {
      (localVideoRef.current.srcObject as MediaStream).getTracks().forEach(track => track.stop());
    }
    router.push('/');
  };

  const removeParticipant = (participantId: string) => {
    setParticipants((prevParticipants) => 
      prevParticipants.filter(participant => participant.id !== participantId)
    );
  };

  return (
    <div className="flex h-screen bg-gradient-to-br from-gray-800 to-black">
      {/* Menu Lateral */}
      <div className="flex flex-col items-center justify-center">
        <div className="w-[70px] bg-gray-400/20 backdrop-blur-lg rounded-e-xl shadow-lg shadow-black/40 border-t border-white/10 flex flex-col items-center py-4 space-y-6">
          <button
            onClick={toggleVideo}
            className={`p-2 rounded-full ${isVideoEnabled ? 'text-emerald-500 hover:text-emerald-600' : 'text-red-500 hover:text-red-600'}`}
          >
            {isVideoEnabled ? <Video /> : <VideoOff />}
          </button>

        <button
          onClick={toggleAudio}
          className={`p-2 rounded-full ${isAudioEnabled ? 'text-emerald-500 hover:text-emerald-600' : 'text-red-500 hover:text-red-600'}`}
        >
          {isAudioEnabled ? <Mic /> : <MicOff />}
        </button>

        <button
          onClick={toggleTabAudio}
          className={`p-2 rounded-full ${!isTabMuted ? 'text-emerald-500 hover:text-emerald-600' : 'text-red-500 hover:text-red-600'}`}
        >
          {!isTabMuted ? <Volume2 /> : <VolumeX />}
        </button>

        <button
          onClick={() => {
            setIsScreenSharing(!isScreenSharing);
            shareScreen()
          }}
          className={`p-2 rounded-full ${isScreenSharing ? 'text-emerald-500 hover:text-emerald-600' : 'text-red-500 hover:text-red-600'}`}
        >
          <Monitor />
        </button>

        <button
          onClick={leaveMeeting}
          className="p-2 rounded-full text-red-500 hover:text-red-600 mt-auto"
        >
          <LogOut />
        </button>
      </div>
      </div>
      {/* Área Principal */}
      <div className="flex-1 p-4">
        <div className="grid auto-rows-fr gap-4 h-full">
          {/* Primeira linha - sempre presente */}
          <div className={`grid ${getGridLayout(Math.min(participants.length, 4))} gap-4`}>
            {participants.map((participant) => (
              <div key={participant.id} className="relative h-[calc(100vh-2rem)]">
                {participant.id === 'local' ? (
                  <>
                    {!isVideoEnabled && (
                      <div className="absolute inset-0 flex items-center justify-center bg-zinc-900 rounded-lg text-white text-2xl font-semibold z-10">
                        {participant.name}
                      </div>
                    )}
                    <video
                      ref={localVideoRef}
                      autoPlay
                      playsInline
                      muted
                      className={`w-full h-full object-cover shadow-lg shadow-black/40 bg-zinc-900/50 rounded-lg ${
                        !isVideoEnabled ? 'invisible' : ''
                      }`}
                    />
                  </>
                ) : (
                  <video
                    ref={remoteVideoRef}
                    autoPlay
                    playsInline
                    className="w-full h-full object-cover shadow-lg shadow-black/40 bg-zinc-900/50 rounded-lg"
                  />
                )}
              </div>
            ))}
          </div>

          {/* Segunda linha - aparece quando há mais de 4 participantes */}
          {participants.length > 4 && (
            <div className={`grid ${getGridLayout(participants.length - 4)} gap-4`}>
              {participants.slice(4).map((participant) => (
                <div key={participant.id} className="relative">
                  <video
                    autoPlay
                    playsInline
                    className="w-full h-full object-cover shadow-lg shadow-black/40 bg-zinc-900/50 rounded-lg"
                  />
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
      <div className="left-menu">
        {participants.map((participant) => (
          <div key={participant.id} className="participant-item flex justify-between items-center p-2">
            <span>{participant.name}</span>
            <button
              onClick={() => removeParticipant(participant.id)}
              className="bg-red-500 text-white px-2 py-1 rounded hover:bg-red-600"
            >
              Remover
            </button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Meeting;
