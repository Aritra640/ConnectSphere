import { useState } from "react";
import { X } from "lucide-react";

interface SigninSignupModalProps {
  ModalState: "Signin" | "Signup";
  onClose: () => void;
}

export function SigninSignupModal({ ModalState, onClose }: SigninSignupModalProps) {
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (ModalState === "Signin") {
      console.log("Sign in with:", email);
    } else {
      console.log("Sign up with:", { username, email, password });
    }
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm">
      <div className="bg-slate-900 border border-slate-800 rounded-2xl shadow-xl w-[90%] max-w-sm p-6 relative">
        {/* Close Button */}
        <button
          onClick={onClose}
          className="absolute top-3 right-3 text-slate-400 hover:text-white"
        >
          <X size={20} />
        </button>

        {/* Header */}
        <h2 className="text-xl font-semibold text-center mb-6 text-white">
          {ModalState === "Signin" ? "Sign In" : "Create an Account"}
        </h2>

        {/* Form */}
        <form onSubmit={handleSubmit} className="space-y-4">
          {ModalState === "Signup" && (
            <input
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full px-4 py-2 rounded-lg bg-slate-800 border border-slate-700 text-white placeholder-slate-400"
              required
            />
          )}

          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full px-4 py-2 rounded-lg bg-slate-800 border border-slate-700 text-white placeholder-slate-400"
            required
          />

          {ModalState === "Signup" && (
            <input
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-4 py-2 rounded-lg bg-slate-800 border border-slate-700 text-white placeholder-slate-400"
              required
            />
          )}

          <button
            type="submit"
            className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-2 rounded-xl transition font-medium"
          >
            {ModalState === "Signin" ? "Sign In" : "Sign Up"}
          </button>
        </form>
      </div>
    </div>
  );
}

