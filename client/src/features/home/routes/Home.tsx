import React from 'react';
import { useUser } from "@/lib/auth"

export const Home = () => {
    const { data: user } = useUser();

    return (
        <div>Hello {user && user.username}!</div>
    )
}