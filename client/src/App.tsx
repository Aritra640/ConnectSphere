import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { HomePage } from "./pages/HomePage";
import { RecoilRoot } from "recoil";
import { DashboardPage } from "./pages/DashboardPage";
import { ChatDashboard } from "./pages/EXPage";

export default function App() {
  return (
    <RecoilRoot>
      <Router>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/dashboard" element={<DashboardPage />} />
          <Route path="/ex" element={<ChatDashboard />} />
        </Routes>
      </Router>
    </RecoilRoot>
  );
}
