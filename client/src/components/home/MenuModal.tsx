import { useRecoilState } from "recoil";
import { X } from "lucide-react";
import { FaGithub } from "react-icons/fa";
import clsx from "clsx";
import { HomeMenuAtom } from "@/store/atoms/homeMenu_atom";

export function MenuModal() {
  const [menuOpen, setMenuOpen] = useRecoilState(HomeMenuAtom);

  if (!menuOpen) return null;

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
          onClick={() => setMenuOpen(false)}
          className="absolute top-3 right-3 text-slate-400 hover:text-white cursor-pointer"
        >
          <X size={20} />
        </button>

        {/* Menu Items */}
        <div className="space-y-5 text-center mt-2 pt-2">
          <button className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-2 rounded-xl transition font-medium cursor-pointer">
            Get Started
          </button>

          <button className="w-full border border-indigo-500 hover:border-white text-indigo-400 hover:text-white py-2 rounded-xl transition font-medium cursor-pointer">
            Login
          </button>

          <a
            href="https://github.com/Aritra640/ConnectSphere"
            target="_blank"
            rel="noopener noreferrer"
            className="w-full flex justify-center items-center gap-2 border border-slate-700 hover:bg-slate-800 text-slate-300 hover:text-white py-2 rounded-xl transition font-medium"
          >
            <FaGithub size={18} />
            Contribute
          </a>

          <button className="w-full text-slate-400 hover:text-white underline text-sm">
            Contact
          </button>
        </div>
      </div>
    </div>
  );
}

