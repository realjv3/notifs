import { useEffect, useState } from "react";
import axios from "axios";
import { Nav } from "./Nav";
import { API_URL } from "./App";

export const Notifications = () => {

    const [notifications, setNotifications] = useState([]);
    const [error, setError] = useState()

    const getNotifications = async () => {
        setError(null)

        try {
            const resp = await axios.get(API_URL + '/notifications', {
                headers: {'Authorization': 'Bearer ' + sessionStorage.getItem('token')},
            });

            setNotifications(resp?.data?.notifications ?? []);
        } catch (error) {
            setError(error?.response?.statusText);
        }
    }

    useEffect(() => {
        getNotifications()
    }, [])

    return (
        <>
            <Nav pref={notifications?.[0]?.Type} />

            <div style={{margin: 'var(--space-500)'}}>
                <h3 className={'h1'}>Notifications</h3>

                <div className='tr'>
                    <div className='td' style={{fontWeight: 'bold'}}>
                        Name
                    </div>

                    <div className='td' style={{fontWeight: 'bold', textAlign: 'center'}}>
                        Type
                    </div>

                    <div className='td' style={{fontWeight: 'bold', textAlign: 'right'}}>
                        Last Edited
                    </div>
                </div>

                <p className={'text-s'} style={{color: 'var(--palette-red-500)'}}>{error}</p>

                {notifications.map((notification) => (
                    <div className='tr' key={notification.ID}>
                        <div className='td'>
                            <span style={{fontWeight: "bold"}}>
                                {notification.Title}
                            </span>
                            <br />
                            <p>{notification.Description}</p><br />
                        </div>

                        <div className='td' style={{textAlign: 'center'}}>
                            {notification.Type}
                        </div>

                        {/* TODO format datetime */}
                        <div className='td' style={{textAlign: 'right'}}>
                            {notification.CreatedAt}
                        </div>
                    </div>
                ))}
            </div>
        </>
    );
}
