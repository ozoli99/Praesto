import { useState, useEffect } from "react";

// Stub hook for interacting with Auth0
const useAuth = () => {
    const [user, setUser] = useState(null);

    useEffect(() => {
        // Initialize auth state from Auth0
    }, []);

    return { user, setUser };
};

export default useAuth;
