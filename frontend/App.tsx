import { BrowserRouter } from "react-router-dom";
import Routes from "./Routes";

export default function App() {
  // Any .tsx or .jsx files in /pages will become a route
  const pages = import.meta.glob("./pages/**/!(*.test.[jt]sx)*.([jt]sx)", {
    eager: true,
  });

  return (
    <BrowserRouter>
      <Routes pages={pages} />
    </BrowserRouter>
  );
}
