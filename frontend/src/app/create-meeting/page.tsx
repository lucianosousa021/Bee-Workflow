'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

const CreateMeeting = () => {
  const router = useRouter();
  const [meetingName, setMeetingName] = useState('');

  const handleCreateMeeting = () => {
    const roomId = Math.random().toString(36).substring(7);
    router.push(`/join-meeting/${roomId}`);
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gradient-to-br from-gray-800 to-black">
      <div className="p-8 bg-gray-900 backdrop-blur-md rounded-xl shadow-2xl shadow-black border border-gray-800/30">
        <h1 className="text-3xl font-bold mb-6 text-gray-100 text-center">Criar Nova Reunião</h1>
        <input
          type="text"
          placeholder="Nome da Reunião"
          className="w-full p-3 mb-6 bg-gray-900/50 border border-gray-700 rounded-lg 
                   text-gray-100 placeholder-gray-400
                   focus:ring-2 focus:ring-gray-600 focus:border-gray-600 
                   outline-none transition-all duration-200"
          value={meetingName}
          onChange={(e) => setMeetingName(e.target.value)}
        />
        <button
          onClick={handleCreateMeeting}
          className="w-full bg-gradient-to-r from-gray-800 to-black text-gray-100 
                   p-3 rounded-lg hover:from-gray-900 hover:to-black 
                   transform hover:scale-[1.02] transition-all duration-200 
                   font-semibold shadow-md hover:shadow-gray-700/50"
        >
          Criar Reunião
        </button>
      </div>
    </div>
  );
};

export default CreateMeeting; 