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
import Link from 'next/link';
type Devices = {
    audio: MediaDeviceInfo[];
    video: MediaDeviceInfo[];
    mic: MediaDeviceInfo[];
};

export default function Home() {
    return (
        <div>
            <Link href="/meeting">Ir para /meeting</Link>
        </div>
    );
}
