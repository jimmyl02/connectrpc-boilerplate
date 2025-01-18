import {
  createConnectQueryKey,
  useMutation,
  useQuery,
  useTransport,
} from "@connectrpc/connect-query";
import { serviceA } from "@repo/gen-api";
import { useQueryClient } from "@tanstack/react-query";
import "./App.css";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";

function App() {
  // get the query client
  const queryClient = useQueryClient();
  const transport = useTransport();

  // create the query and mutation hooks
  const { data } = useQuery(serviceA.getCount);
  const updateCount = useMutation(serviceA.newCount, {
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: createConnectQueryKey({
          schema: serviceA.getCount,
          transport: transport,
          cardinality: "finite",
        }),
      });
    },
  });

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => updateCount.mutate({ amount: 1 })}>
          count is {data && data.count ? data.count : 0}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
