import { useRecoilState, useSetRecoilState } from "recoil";
import { X } from "lucide-react";
import clsx from "clsx";
import { AuthModalAtom } from "@/store/atoms/authmodal_atom";
import { AuthStateAtom } from "@/store/atoms/authstate_atom";
import { OTPModalAtom } from "@/store/atoms/otpmodal_atom";
import { useNavigate } from "react-router-dom";

export function AuthModal() {
  const [modalOpen, setModalOpen] = useRecoilState(AuthModalAtom);
  const [authState, setAuthState] = useRecoilState(AuthStateAtom);
  const navigate = useNavigate();
  const setOtpModal = useSetRecoilState(OTPModalAtom);
  const setAuthModal = useSetRecoilState(AuthModalAtom);

  function SigninSignupProcess() {

    if (authState == "Signin") {

      //do all signin stuff
      navigate("/dashboard");
    }


    if (authState == "Signup") {
      //send signup process to server

      setOtpModal(true);
      setAuthModal(false);
    }
  }

  if (!modalOpen) return null;

  return (
    <div
      className={clsx(
        "fixed inset-0 z-50 flex items-center justify-center",
        "bg-black/60 backdrop-blur-sm"
      )}
    >
      <div className="bg-slate-900 border border-slate-800 rounded-2xl shadow-xl w-96 p-6 relative text-white">
        {/* Close Button */}
        <button
          onClick={() => setModalOpen(false)}
          className="absolute top-3 right-3 text-slate-400 hover:text-white cursor-pointer"
        >
          <X size={20} />
        </button>

        <h2 className="text-2xl font-semibold text-center mb-6">
          {authState === "Signin" ? "Sign In to ConnectSphere" : "Create an Account"}
        </h2>

        <form className="space-y-4">
          {authState === "Signup" && (
            <input
              type="text"
              placeholder="Username"
              className="w-full bg-slate-800 text-white px-4 py-2 rounded-xl border border-slate-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 cursor-pointer"
            />
          )}

          <input
            type="email"
            placeholder="Email"
            className="w-full bg-slate-800 text-white px-4 py-2 rounded-xl border border-slate-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 cursor-pointer"
          />

          <input
            type="password"
            placeholder="Password"
            className="w-full bg-slate-800 text-white px-4 py-2 rounded-xl border border-slate-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 cursor-pointer"
          />

          <button
            onClick={SigninSignupProcess}
            type="submit"
            className="mt-2 w-full bg-indigo-600 hover:bg-indigo-700 text-white py-2 rounded-xl font-medium transition cursor-pointer"
          >
            {authState === "Signin" ? "Sign In" : "Sign Up"}
          </button>
        </form>

        <div className="mt-4 text-sm text-center text-slate-400">
          {authState === "Signin" ? (
            <>
              Don’t have an account?{" "}
              <button
                onClick={() => setAuthState("Signup")}
                className="text-indigo-400 hover:underline"
              >
                Sign up
              </button>
            </>
          ) : (
            <>
              Already have an account?{" "}
              <button
                onClick={() => setAuthState("Signin")}
                className="text-indigo-400 hover:underline"
              >
                Sign in
              </button>
            </>
          )}
        </div>
      </div>
    </div>
  );
}

