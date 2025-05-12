import { useRecoilState } from "recoil";
import { OTPModalAtom } from "@/store/atoms/otpmodal_atom";
import { X } from "lucide-react";
import clsx from "clsx";
import { useState } from "react";

export function OTPModal() {
  const [otpModal, setOtpModal] = useRecoilState(OTPModalAtom);
  const [otp, setOtp] = useState("");

  if (!otpModal) return null;

  function handleVerify() {
    // Add actual OTP verification logic here
    console.log("Entered OTP: ", otp);
    setOtpModal(false);
  }

  return (
    <div
      className={clsx(
        "fixed inset-0 z-50 flex items-center justify-center",
        "bg-black/60 backdrop-blur-sm"
      )}
    >
      <div className="bg-slate-900 border border-slate-800 rounded-2xl shadow-xl w-80 p-6 relative">
        {/* Close Button */}
        <button
          onClick={() => setOtpModal(false)}
          className="absolute top-3 right-3 text-slate-400 hover:text-white cursor-pointer"
        >
          <X size={20} />
        </button>

        <h2 className="text-xl text-white font-semibold mb-4 text-center">
          Enter OTP
        </h2>

        <input
          type="text"
          value={otp}
          onChange={(e) => setOtp(e.target.value)}
          placeholder="Enter OTP"
          className="w-full px-4 py-2 rounded-lg bg-slate-800 text-white border border-slate-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 mb-4 cursor-pointer"
        />

        <button
          onClick={handleVerify}
          className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-2 rounded-xl transition font-medium cursor-pointer"
        >
          Verify OTP
        </button>
      </div>
    </div>
  );
}

