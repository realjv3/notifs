import { useState } from "react";
import axios from "axios";
import { API_URL } from "./App";

export const Login = () => {

    const [email, setEmail] = useState();
    const [error, setError] = useState();

    const login = async () => {
        setError(null);

        try {
            const resp = await axios({
                method: 'post',
                url: API_URL + '/login',
                headers: {"Content-Type": "application/x-www-form-urlencoded"},
                data: {email}
            });

            sessionStorage.setItem('token', resp?.data?.token);
            sessionStorage.setItem('email', email);

            window.location.reload();
        } catch (error) {
            setError(error?.response?.statusText);
        }
    };

    return (
        <fieldset>
            <h3 className={'h3'}>Welcome Back!</h3>

            <p>Sign in to access your notifications</p>

            <label htmlFor="email">Email</label>
            <input
                type="email"
                id="email"
                style={{width: '80%', marginRight:'var(--space-175)'}}
                onChange={(v) => setEmail(v.currentTarget.value)}
            />
            <button
                type="submit"
                className="dark-btn"
                onClick={login}
            >
                Sign In
            </button>

            <p
                className={'text-s'}
                style={{color: 'var(--palette-red-500)', margin: 'var(--space-050)'}}
            >
                {error}
            </p>
        </fieldset>
    );
}
