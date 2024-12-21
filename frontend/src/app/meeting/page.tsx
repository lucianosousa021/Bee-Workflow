'use client';

import React, { useEffect, useRef, useState } from "react";

const VideoCall: React.FC = () => {
  const localVideoRef = useRef<HTMLVideoElement>(null);
  const remoteVideoRef = useRef<HTMLVideoElement>(null);
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const peerConnection = useRef<RTCPeerConnection | null>(null);

  useEffect(() => {
    const initWebRTC = async () => {
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
          { urls: "turn:yourdomain.com", username: "user", credential: "password" },
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
    };

    initWebRTC();

    const ws = new WebSocket("ws://localhost:8080/ws?id=user1&room=room1");
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

  const createOffer = async () => {
    if (peerConnection.current && socket) {
      const offer = await peerConnection.current.createOffer();
      await peerConnection.current.setLocalDescription(offer);
      socket.send(
        JSON.stringify({
          type: "offer",
          payload: offer,
        })
      );
    }
  };

  return (
    <div>
      <video ref={localVideoRef} autoPlay muted playsInline />
      <video ref={remoteVideoRef} autoPlay playsInline />
      <button onClick={createOffer}>Start Call</button>
    </div>
  );
};

export default VideoCall;
