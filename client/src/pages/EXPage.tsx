import {
  Home,
  MessageCircle,
  Users,
  Settings,
  Circle,
} from "lucide-react";
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar";
import { useRecoilValue } from "recoil";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { MainThemeIcon } from "@/components/global/MainThemeIcon";

const messages = [
  { name: "Anna", message: "Sure! That works for me.", time: "10:40 AM", img: "/users/anna.jpg" },
  { name: "James", message: "See you there!", time: "9:20 AM", img: "/users/james.jpg" },
  { name: "Amanda", message: "I sent you the files.", time: "Yesterday", img: "/users/amanda.jpg" },
  { name: "Michael", message: "Got it, thanks!", time: "Yesterday", img: "/users/michael.jpg" },
  { name: "Rachel", message: "What do you think?", time: "Sunday", img: "/users/rachel.jpg" },
  { name: "Jason", message: "Let’s do it.", time: "Sunday", img: "/users/jason.jpg" },
];

const activeUsers = ["Daniel", "Alicia", "Steven"];

export function ChatDashboard() {
  const theme = useRecoilValue(MainThemeAtom);

  return (
    <div
      className={`flex h-screen transition-colors duration-300 ${
        theme === "Dark" ? "bg-gray-900 text-white" : "bg-gray-100 text-black"
      }`}
    >
      {/* Sidebar */}
      <aside
        className={`w-20 md:w-64 ${
          theme === "Dark" ? "bg-gray-800" : "bg-purple-700"
        } text-white flex flex-col justify-between p-4 md:p-6 transition-all duration-300`}
      >
        <div>
          <h2 className="text-xl md:text-2xl font-bold mb-6 md:mb-10 hidden md:block">Dashboard</h2>
          <nav className="space-y-4">
            <NavItem icon={<Home />} text="Home" />
            <NavItem icon={<MessageCircle />} text="Messages" />
            <NavItem icon={<Users />} text="Groups" />
            <NavItem icon={<Settings />} text="Settings" />
          </nav>
        </div>

        <div className="mt-10">
          <div
            className={`flex items-center justify-center md:justify-start space-x-0 md:space-x-3 p-2 md:p-3 rounded-lg ${
              theme === "Dark" ? "bg-gray-700" : "bg-purple-800"
            }`}
          >
            <Avatar>
              <AvatarImage src="/users/johndoe.jpg" alt="John Doe" />
              <AvatarFallback>JD</AvatarFallback>
            </Avatar>
            <span className="font-medium hidden md:inline">John Doe</span>
          </div>
        </div>
      </aside>

      {/* Main Content */}
      <main className="flex-1 p-4 md:p-6 space-y-6 overflow-y-auto">
        {/* Header */}
        <div className="flex justify-between items-center">
          <h1 className="text-2xl md:text-3xl font-bold">Chats</h1>
          <MainThemeIcon />
        </div>

        {/* Stats */}
        <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
          <StatCard label="Direct Messages" value="12" />
          <StatCard label="Group Messages" value="5" />
          <StatCard label="Time Spent" value="4h 12m" />
        </div>

        {/* Direct Messages */}
        <section
          className={`rounded-xl shadow p-4 md:p-5 ${
            theme === "Dark" ? "bg-gray-800" : "bg-white"
          }`}
        >
          <h2 className="text-lg md:text-xl font-semibold mb-4">Direct Messages</h2>
          <div className="space-y-4">
            {messages.map((msg) => (
              <div key={msg.name} className="flex justify-between items-start">
                <div className="flex items-center space-x-3">
                  <Avatar>
                    <AvatarImage src={msg.img} alt={msg.name} />
                    <AvatarFallback>
                      {msg.name
                        .split(" ")
                        .map((part) => part[0])
                        .join("")
                        .toUpperCase()}
                    </AvatarFallback>
                  </Avatar>
                  <div>
                    <h3 className="font-semibold">{msg.name}</h3>
                    <p className="text-sm text-gray-600 dark:text-gray-400">{msg.message}</p>
                  </div>
                </div>
                <span className="text-sm text-gray-500 dark:text-gray-400">{msg.time}</span>
              </div>
            ))}
          </div>
        </section>

        {/* Active Now */}
        <section
          className={`rounded-xl shadow p-4 md:p-5 ${
            theme === "Dark" ? "bg-gray-800" : "bg-white"
          }`}
        >
          <h2 className="text-lg md:text-xl font-semibold mb-4">Active Now</h2>
          <ul className="space-y-3">
            {activeUsers.map((name) => (
              <li key={name} className="flex items-center justify-between">
                <div className="flex items-center space-x-3">
                  <Avatar>
                    <AvatarImage src={`/users/${name.toLowerCase()}.jpg`} alt={name} />
                    <AvatarFallback>
                      {name
                        .split(" ")
                        .map((part) => part[0])
                        .join("")
                        .toUpperCase()}
                    </AvatarFallback>
                  </Avatar>
                  <span>{name}</span>
                </div>
                <Circle className="w-3 h-3 text-green-500 fill-green-500" />
              </li>
            ))}
          </ul>
        </section>
      </main>
    </div>
  );
}

function NavItem({ icon, text }: { icon: React.ReactNode; text: string }) {
  return (
    <div className="flex items-center justify-center md:justify-start space-x-0 md:space-x-3 cursor-pointer hover:bg-purple-600 dark:hover:bg-gray-700 p-2 md:p-3 rounded-lg">
      {icon}
      <span className="hidden md:inline">{text}</span>
    </div>
  );
}

function StatCard({ label, value }: { label: string; value: string }) {
  return (
    <div className="rounded-xl shadow p-4 text-center bg-white dark:bg-gray-800">
      <p className="text-2xl font-bold">{value}</p>
      <p className="text-gray-500 dark:text-gray-400">{label}</p>
    </div>
  );
}

