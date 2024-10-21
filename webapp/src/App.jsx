import { Login } from "./Login";
import { Notifications } from "./Notifications";

export const API_URL = "http://localhost:8080";

export const App = () => {
    if (sessionStorage.getItem('token')) {
        return <Notifications />;
    }

    return <Login />
}
