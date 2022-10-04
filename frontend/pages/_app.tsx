import "../styles/style.scss";
import type { AppProps } from "next/app";
import { Provider } from "react-redux";

import { store } from "app/store";
import { RouteGuard } from "../components/AuthGuard";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <RouteGuard>
      <Provider store={store}>
        <Component {...pageProps} />
      </Provider>
    </RouteGuard>
  );
}

export default MyApp;
