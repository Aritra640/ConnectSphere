import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { HomePage } from "./pages/HomePage";
import { RecoilRoot } from "recoil";

export default function App() {

  return (
    <RecoilRoot>
    <Router>

      <Routes>

        <Route path="/" element={<HomePage />} />
        

      </Routes>

    </Router>
    </RecoilRoot>
  );
}
