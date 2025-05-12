import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { HomePage } from "./pages/HomePage";
import { RecoilRoot } from "recoil";
import { DashboardPage } from "./pages/DashboardPage";

export default function App() {

  return (
    <RecoilRoot>
    <Router>

      <Routes>

        <Route path="/" element={<HomePage />} />
        <Route path="/dashboard" element={<DashboardPage />} />
        
      </Routes>

    </Router>
    </RecoilRoot>
  );
}
