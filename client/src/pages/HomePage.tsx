import { useRecoilState } from "recoil";
import { HomeMenuButton } from "@/components/global/menubutton";
import { ConnectSphereIcon } from "@/icons/connectsphere_icon";
import { HomeMenuAtom } from "@/store/atoms/homeMenu_atom";
import { MenuModal } from "@/components/home/MenuModal";

export function HomePage() {
  const [menuOpen, setMenuOpen] = useRecoilState(HomeMenuAtom);

  console.log("home menu status: " , menuOpen);

  return (
    <div className="bg-slate-950 min-h-screen w-screen text-white font-sans overflow-hidden">
      {/* Top Navigation */}
      <header className="sticky top-0 z-50 bg-slate-950 flex justify-between items-center px-4 py-3 border-b border-slate-800">
        {/* Logo + Title */}
        <div className="flex items-center space-x-3">
          <ConnectSphereIcon />
          <span className="text-2xl md:text-3xl font-mono text-white">ConnectSphere</span>
        </div>

        {/* Buttons + Menu */}
        <div className="flex items-center space-x-4">
          {/* Hidden on small screens */}
          <div className="hidden md:flex items-center space-x-4">
            <button className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-xl font-medium transition cursor-pointer">
              Get Started
            </button>
            <button className="border border-indigo-600 text-indigo-400 hover:text-white hover:border-white px-4 py-2 rounded-xl font-medium transition cursor-pointer">
              Sign In
            </button>
          </div>

          {/* Menu Button */}
          <button onClick={() => setMenuOpen(true)}>
            <HomeMenuButton />
          </button>
        </div>
      </header>

      {/* Modal */}
      <MenuModal />

      {/* Main Content */}
      <main className="p-6">
        <section className="max-w-3xl mx-auto space-y-10">
          <div className="space-y-4">
            <h1 className="text-4xl md:text-5xl font-bold text-white">
              Connect. Share. Belong.
            </h1>
            <p className="text-slate-300 text-lg md:text-xl">
              ConnectSphere is your digital space to build meaningful relationships—whether it’s with friends, teammates, or your community.
              Enjoy real-time conversations, vibrant group chats, and private messaging all in one place.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mt-4">
            <div className="bg-slate-900 rounded-xl p-5 shadow-lg border border-slate-800 hover:shadow-indigo-800/30 transition">
              <h2 className="text-xl md:text-2xl font-semibold mb-2">🔒 Private Messaging</h2>
              <p className="text-slate-400 text-sm md:text-base">
                Secure, one-on-one conversations with full end-to-end encryption. Keep your chats private and personal.
              </p>
            </div>

            <div className="bg-slate-900 rounded-xl p-5 shadow-lg border border-slate-800 hover:shadow-indigo-800/30 transition">
              <h2 className="text-xl md:text-2xl font-semibold mb-2">👥 Group Channels</h2>
              <p className="text-slate-400 text-sm md:text-base">
                Build thriving communities or tight-knit teams. Group channels let you collaborate and chill, your way.
              </p>
            </div>

            <div className="bg-slate-900 rounded-xl p-5 shadow-lg border border-slate-800 hover:shadow-indigo-800/30 transition">
              <h2 className="text-xl md:text-2xl font-semibold mb-2">📣 Instant Notifications</h2>
              <p className="text-slate-400 text-sm md:text-base">
                Stay in the loop with real-time alerts for mentions, messages, invites, and more—never miss a beat.
              </p>
            </div>
          </div>
        </section>
      </main>
    </div>
  );
}

