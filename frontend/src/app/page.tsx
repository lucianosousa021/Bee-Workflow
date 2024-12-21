// pages/index.js
'use client';

import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

export default function Home() {
    const router = useRouter();

    useEffect(() => {
        router.push('/create-meeting');
    }, [router]);

    return (
        <div>
            <Link href="/create-meeting">Ir para /create-meeting</Link>
        </div>
    );
}
