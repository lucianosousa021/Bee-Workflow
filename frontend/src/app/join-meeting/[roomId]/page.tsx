'use client';

import { useState } from 'react';
import { useRouter, useParams } from 'next/navigation';

const JoinMeeting = () => {
  const router = useRouter();
  const params = useParams();
  const [userName, setUserName] = useState('');

  const handleJoinMeeting = () => {
    if (userName.trim()) {
      router.push(`/meeting?room=${params.roomId}&user=${userName}`);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gradient-to-br from-gray-800 to-black">
      <div className="p-8 bg-gray-900 backdrop-blur-md rounded-xl shadow-2xl shadow-black border border-gray-800/30">
        <h1 className="text-2xl font-bold mb-6 text-gray-100 text-center">Entrar na Reunião</h1>
        <p className="mb-4">ID da Sala: {params.roomId}</p>
        <input
          type="text"
          placeholder="Seu Nome"
          className="w-full p-3 mb-6 bg-gray-900/50 border border-gray-700 rounded-lg 
                   text-gray-100 placeholder-gray-400
                   focus:ring-2 focus:ring-gray-600 focus:border-gray-600 
                   outline-none transition-all duration-200"
          value={userName}
          onChange={(e) => setUserName(e.target.value)}
        />
        <button
          onClick={handleJoinMeeting}
          className="w-full bg-gradient-to-r from-gray-800 to-black text-gray-100 
                   p-3 rounded-lg hover:from-gray-900 hover:to-black 
                   transform hover:scale-[1.02] transition-all duration-200 
                   font-semibold shadow-md hover:shadow-gray-700/50"
        >
          Entrar na Reunião
        </button>
      </div>
    </div>
  );
};

export default JoinMeeting;