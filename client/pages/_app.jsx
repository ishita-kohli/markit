import "../styles/globals.css";
import AuthContextProvider from "../modules/auth_provider";
import WebSocketProvider from "../modules/websocket_provider";
import "../axios.config";

export default function App({ Component, pageProps }) {
  return (
    <>
      <AuthContextProvider>
        <WebSocketProvider>
          <div className="flex flex-col h-full min-h-screen font-sans md:flex-row">
            <Component {...pageProps} />
          </div>
        </WebSocketProvider>
      </AuthContextProvider>
    </>
  );
}
