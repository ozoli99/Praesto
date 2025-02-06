import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import api from "../api/axiosInstance.js";

const Profile = () => {
    const { user, error, isAuthenticated } = useAuth0();
    const [profileData, setProfileData] = React.useState(null);

    React.useEffect(() => {
        const fetchProfile = async () => {
            try {
                const response = await api.get("/profile");
                setProfileData(response.data);
            } catch (err) {
                console.error(
                    "API error:",
                    err.response ? err.response.data : err
                );
            }
        };
        if (isAuthenticated) {
            fetchProfile();
        }
    }, [isAuthenticated]);

    if (error) return <div>Authentication Error: {error.message}</div>;

    return (
        <div>
            <h2>Profile</h2>
            {profileData ? (
                <pre>{JSON.stringify(profileData, null, 2)}</pre>
            ) : (
                <div>Loading...</div>
            )}
        </div>
    );
};

export default Profile;
