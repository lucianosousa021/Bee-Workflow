// pages/index.js
'use client';

import { useState, useEffect, useRef } from 'react';
import { createPortal } from 'react-dom';
import { 
    Video, 
    VideoOff, 
    Mic, 
    MicOff, 
    Monitor, 
    Settings, 
    LogOut, 
    LayoutGrid 
} from 'lucide-react';

type Devices = {
    audio: MediaDeviceInfo[];
    video: MediaDeviceInfo[];
    mic: MediaDeviceInfo[];
};

export default function Home() {
    const [isVideoOn, setIsVideoOn] = useState(true);
    const [isAudioOn, setIsAudioOn] = useState(true);
    const [isScreenSharing, setIsScreenSharing] = useState(false);
    const [modalOpen, setModalOpen] = useState(false);
    const [devices, setDevices] = useState<Devices>({ audio: [], video: [], mic: [] });
    const videoRef = useRef<HTMLVideoElement | null>(null);
    const streamRef = useRef<MediaStream | null>(null);
    let videoSender: RTCRtpSender | null = null;
    let peerConnection: RTCPeerConnection | null = null;
    const [isExpanded, setIsExpanded] = useState(false);

    useEffect(() => {
        const initMedia = async () => {
            const stream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });
            if (videoRef.current) {
                videoRef.current.srcObject = stream;
            }
            streamRef.current = stream;
            const deviceList = await navigator.mediaDevices.enumerateDevices();

            setDevices({
                audio: deviceList.filter(device => device.kind === 'audiooutput'),
                video: deviceList.filter(device => device.kind === 'videoinput'),
                mic: deviceList.filter(device => device.kind === 'audioinput'),
            });
        };
        initMedia();

        // Inicialize a conexão WebRTC
        peerConnection = new RTCPeerConnection();

        return () => {
            streamRef.current?.getTracks().forEach(track => track.stop());
            peerConnection?.close();
            peerConnection = null;
        };
    }, []);

    const toggleVideo = () => {
        const videoTrack = streamRef.current?.getVideoTracks()[0];
        if (videoTrack) {
            videoTrack.enabled = !isVideoOn;
            setIsVideoOn(!isVideoOn);
        }
    };

    const toggleAudio = () => {
        const audioTrack = streamRef.current?.getAudioTracks()[0];
        if (audioTrack) {
            audioTrack.enabled = !isAudioOn;
            setIsAudioOn(!isAudioOn);
        }
    };

    const addVideoTrackToConnection = (track: MediaStreamTrack) => {
        if (peerConnection && streamRef.current) {
            videoSender = peerConnection.addTrack(track, streamRef.current);
        }
    };

    const shareScreen = async () => {
        if (!isScreenSharing && streamRef.current) {
            const screenStream = await navigator.mediaDevices.getDisplayMedia({ video: true });
            const videoTrack = screenStream.getVideoTracks()[0];

            if (videoSender) {
                await videoSender.replaceTrack(videoTrack);
            }

            videoTrack.onended = () => stopScreenSharing();
            setIsScreenSharing(true);
        } else {
            stopScreenSharing();
        }
    };

    const stopScreenSharing = () => {
        if (streamRef.current) {
            const videoTrack = streamRef.current.getVideoTracks()[0];
            videoTrack.stop();
            setIsScreenSharing(false);
        }
    };

    const openModal = () => setModalOpen(true);
    const closeModal = () => setModalOpen(false);

    return (
        <div className="h-screen flex">
            {/* Sidebar */}
            <aside 
                className={`bg-gray-800 text-white p-4 transition-all duration-300 ease-in-out ${
                    isExpanded ? 'w-64' : 'w-16'
                } overflow-hidden`}
                onMouseEnter={() => setIsExpanded(true)}
                onMouseLeave={() => setIsExpanded(false)}
            >
                <div className="space-y-4">
                    <div className={`w-full p-2.5 bg-gray-700 hover:bg-gray-600 rounded-lg transition-colors flex items-center ${
                        isExpanded ? 'px-4' : 'justify-center'
                    }`}>
                        <button 
                            onClick={toggleVideo}
                            className="flex items-center w-full"
                        >
                            <div className="flex-shrink-0">
                                {isVideoOn ? <Video size={20} /> : <VideoOff size={20} />}
                            </div>
                            {isExpanded && (
                                <span className="ml-3 whitespace-nowrap overflow-hidden text-ellipsis">
                                    {isVideoOn ? 'Desligar Vídeo' : 'Ligar Vídeo'}
                                </span>
                            )}
                        </button>
                    </div>

                    <div className={`w-full p-2.5 bg-gray-700 hover:bg-gray-600 rounded-lg transition-colors flex items-center ${
                        isExpanded ? 'px-4' : 'justify-center'
                    }`}>
                        <button 
                            onClick={toggleAudio}
                            className="flex items-center w-full"
                        >
                            <div className="flex-shrink-0">
                                {isAudioOn ? <Mic size={20} /> : <MicOff size={20} />}
                            </div>
                            {isExpanded && (
                                <span className="ml-3 whitespace-nowrap overflow-hidden text-ellipsis">
                                    {isAudioOn ? 'Desligar Áudio' : 'Ligar Áudio'}
                                </span>
                            )}
                        </button>
                    </div>

                    <div className={`w-full p-2.5 bg-gray-700 hover:bg-gray-600 rounded-lg transition-colors flex items-center ${
                        isExpanded ? 'px-4' : 'justify-center'
                    }`}>
                        <button 
                            onClick={shareScreen}
                            className="flex items-center w-full"
                        >
                            <div className="flex-shrink-0">
                                <Monitor size={20} />
                            </div>
                            {isExpanded && (
                                <span className="ml-3 whitespace-nowrap overflow-hidden text-ellipsis">
                                    {isScreenSharing ? 'Parar Compartilhar' : 'Compartilhar Tela'}
                                </span>
                            )}
                        </button>
                    </div>

                    <div className={`w-full p-2.5 bg-gray-700 hover:bg-gray-600 rounded-lg transition-colors flex items-center ${
                        isExpanded ? 'px-4' : 'justify-center'
                    }`}>
                        <button 
                            onClick={openModal}
                            className="flex items-center w-full"
                        >
                            <div className="flex-shrink-0">
                                <Settings size={20} />
                            </div>
                            {isExpanded && (
                                <span className="ml-3 whitespace-nowrap overflow-hidden text-ellipsis">
                                    Dispositivos
                                </span>
                            )}
                        </button>
                    </div>

                    <div className={`w-full p-2.5 bg-red-600 hover:bg-red-700 rounded-lg transition-colors flex items-center ${
                        isExpanded ? 'px-4' : 'justify-center'
                    }`}>
                        <button 
                            className="flex items-center w-full"
                        >
                            <div className="flex-shrink-0">
                                <LogOut size={20} />
                            </div>
                            {isExpanded && (
                                <span className="ml-3 whitespace-nowrap overflow-hidden text-ellipsis">
                                    Sair
                                </span>
                            )}
                        </button>
                    </div>

                    <div className={`w-full p-2.5 bg-gray-700 hover:bg-gray-600 rounded-lg transition-colors flex items-center ${
                        isExpanded ? 'px-4' : 'justify-center'
                    }`}>
                        <button 
                            className="flex items-center w-full"
                        >
                            <div className="flex-shrink-0">
                                <LayoutGrid size={20} />
                            </div>
                            {isExpanded && (
                                <span className="ml-3 whitespace-nowrap overflow-hidden text-ellipsis">
                                    Layout
                                </span>
                            )}
                        </button>
                    </div>
                </div>
            </aside>

            {/* Video Section */}
            <div className="flex-1 flex items-center justify-center bg-black">
                <video ref={videoRef} autoPlay playsInline className="w-full h-full object-cover" />
            </div>

            {/* Modal */}
            {modalOpen && createPortal(
                <div className="fixed inset-0 bg-gray-900 bg-opacity-50 flex items-center justify-center">
                    <div className="bg-white p-4 rounded shadow">
                        <h2 className="text-xl mb-4">Selecionar Dispositivos</h2>
                        <div>
                            <label>Vídeo:</label>
                            <select className="block w-full mb-2">
                                {devices.video.map(device => (
                                    <option key={device.deviceId} value={device.deviceId}>{device.label}</option>
                                ))}
                            </select>
                        </div>
                        <div>
                            <label>Áudio:</label>
                            <select className="block w-full mb-2">
                                {devices.audio.map(device => (
                                    <option key={device.deviceId} value={device.deviceId}>{device.label}</option>
                                ))}
                            </select>
                        </div>
                        <div>
                            <label>Microfone:</label>
                            <select className="block w-full mb-2">
                                {devices.mic.map(device => (
                                    <option key={device.deviceId} value={device.deviceId}>{device.label}</option>
                                ))}
                            </select>
                        </div>
                        <button className="bg-blue-500 text-white p-2 rounded" onClick={closeModal}>Fechar</button>
                    </div>
                </div>,
                document.body
            )}
        </div>
    );
}
