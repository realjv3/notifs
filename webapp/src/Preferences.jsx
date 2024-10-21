import {useEffect, useState} from "react";
import axios from "axios";
import { API_URL } from "./App";

export const Preferences = ({ isOpen, toggle, currentPref }) => {
    const [preference, setPreference] = useState(currentPref);

    useEffect(() => {
        setPreference(currentPref);
    }, [currentPref, setPreference]);

    const changePref = async () => {
        try {
            const resp = await axios({
                method: 'post',
                url: API_URL + '/preferences',
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                    'Authorization': 'Bearer ' + sessionStorage.getItem('token'),
                },
                data: {preference}
            });

            if (resp) {
                toggle(false);
            }
        } catch (error) {
            alert(error?.response?.statusText)
        }
    };

    return (
        <>
            <div className="overlay" style={{display: isOpen ? 'block' : 'none'}}></div>

            <dialog className="modal" style={{display: isOpen ? 'block' : 'none'}}>
                <div style={{display: "flex", justifyContent: 'space-between'}}>
                    <p className={'h3'}>Preferences</p>
                    <p className="close" onClick={() => toggle(false)}>
                        &times;
                    </p>
                </div>

                <p>What type of notifications do you want to see?</p>

                <div className="checkbox">
                    {/*TODO accommodate multiple preferences*/}
                    <input type="checkbox" id="select-all" />
                    <label htmlFor="select-all">Select all</label>
                </div>

                <hr />

                <div className="checkbox">
                    <input
                        type="checkbox"
                        id="email"
                        checked={preference === 'EMAIL'}
                        onChange={() => setPreference('EMAIL')}
                    />
                    <label htmlFor="email">Email</label>
                </div>

                <div className="checkbox">
                    <input
                        type="checkbox"
                        id="sms"
                        checked={preference === 'SMS'}
                        onChange={() => setPreference('SMS')}
                    />
                    <label htmlFor="sms">SMS</label>
                </div>

                <div className="checkbox">
                    <input
                        type="checkbox"
                        id="push"
                        checked={preference === 'PUSH'}
                        onChange={() => setPreference('PUSH')}
                    />
                    <label htmlFor="push">Push</label>
                </div>

                <div className="modal-actions">
                    <button onClick={() => toggle(false)}>
                        Cancel
                    </button>

                    <button className='dark-btn' onClick={changePref}>
                        Submit
                    </button>
                </div>
            </dialog>
        </>
    );
}