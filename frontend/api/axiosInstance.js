import axios from "axios";
import { getAccessTokenSilently } from "@auth0/auth0-react";

const api = axios.create({
    baseURL: process.env.REACT_APP_API_URL || "http://localhost:8080",
});

export const useAxiosWithAuth = () => {
    const { getAccessTokenSilently } = useAuth0();
    api.interceptors.request.use(async (config) => {
        const token = await getAccessTokenSilently();
        config.headers.Authorization = `Bearer ${token}`;
        return config;
    });
    return api;
};

export default api;
