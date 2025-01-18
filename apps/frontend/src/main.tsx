import { TransportProvider } from "@connectrpc/connect-query";
import { createConnectTransport } from "@connectrpc/connect-web";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import "./index.css";

const finalTransport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

const queryClient = new QueryClient();

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <TransportProvider transport={finalTransport}>
      <QueryClientProvider client={queryClient}>
        <App />
      </QueryClientProvider>
    </TransportProvider>
  </StrictMode>
);
