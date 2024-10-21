import { useState } from "react";
import { Preferences } from "./Preferences";

export const Nav = ({ pref }) => {
    const [showPrefs, setShowPrefs] = useState(false);

    const logout = () => {
        sessionStorage.clear();
        window.location.reload();
    };

    return (
        <>
            <div className='nav'>
                <span className={'text-s'}>Notification System</span>

                <div>
                    <span style={{color: 'var(--palette-grey-300)'}}>
                        {sessionStorage.getItem('email')}
                    </span>

                    <span style={{color: 'var(--palette-grey-300)', margin: '0 7px 0'}}>|</span>

                    <span style={{cursor: 'pointer'}} onClick={() => setShowPrefs(true)}>
                        Preferences
                    </span>

                    <span onClick={logout} style={{cursor: 'pointer'}}>
                        Log out
                    </span>
                </div>
            </div>

            <Preferences isOpen={showPrefs} toggle={setShowPrefs} currentPref={pref} />
        </>
    );
}
