import { MainThemeIcon } from "@/components/global/MainThemeIcon";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { Avatar, AvatarFallback, AvatarImage } from "@radix-ui/react-avatar";
import { Home, MessageCircle, Settings, Users } from "lucide-react";
import { useRecoilValue } from "recoil";

const DashboardTheme = {
  "Bright": "bg-gray-200 text-black",
  "Dark": "bg-gray-950 text-white",
};

const SidebarTheme = {
  "Bright": "bg-purple-800",
  "Dark": "bg-gray-800",
};

const CardThemes = {
  "Bright": "bg-white",
  "Dark": "bg-gray-800",
}

const messages = [
  { name: "Anna", message: "Sure! That works for me.", time: "10:40 AM", img: "/users/anna.jpg" },
  { name: "James", message: "See you there!", time: "9:20 AM", img: "/users/james.jpg" },
  { name: "Amanda", message: "I sent you the files.", time: "Yesterday", img: "/users/amanda.jpg" },
  { name: "Michael", message: "Got it, thanks!", time: "Yesterday", img: "/users/michael.jpg" },
  { name: "Rachel", message: "What do you think?", time: "Sunday", img: "/users/rachel.jpg" },
  { name: "Jason", message: "Let’s do it.", time: "Sunday", img: "/users/jason.jpg" },
];


const groupMessages = [
  { name: "Group1", sender: "user1", message: "Hi hello there", time: "10:17 AM", img: ""},
  { name: "Group2", sender: "user2", message: "Hi hello there", time: "10:17 AM", img: ""},
  { name: "Group1", sender: "user1", message: "Hi hello there", time: "10:17 AM", img: ""},
  { name: "Group1", sender: "user1", message: "Hi hello there", time: "10:17 AM", img: ""},
];


export function DashboardPage() {
  const defaultColorTheme = "flex h-screen transition-colors duration-300";
  const theme = useRecoilValue(MainThemeAtom);

  return (
    <div className={DashboardTheme[theme] + " " + defaultColorTheme}>
      <aside
        className={
          "w-20 md:w-64" +
          " " +
          SidebarTheme[theme] +
          " text-white flex flex-col justify-start gap-14 p-4 md:p-6 transition-all duration-300"
        }
      >

        <div className="">
          <div
            className={`flex items-center justify-center md:justify-start space-x-0 md:space-x-3 p-2 md:p-3 rounded-lg cursor-pointer ${
              theme === "Dark" ? "bg-gray-700 hover:bg-gray-600" : "bg-purple-800 hover:bg-purple-700"
            }`}
          >
            <Avatar>
              <AvatarImage src="/users/johndoe.jpg" alt="John Doe" />
              <AvatarFallback>JD</AvatarFallback>
            </Avatar>
            <span className="font-medium hidden md:inline">John Doe</span>
          </div>
        </div>


        <div>
          <nav className="space-y-4">
            <NavItem icon={<Home />} text="Home" />
            <NavItem icon={<MessageCircle />} text="Messages" />
            <NavItem icon={<Users />} text="Groups" />
            <NavItem icon={<Settings />} text="Settings" />
          </nav>
        </div>
      </aside>


      <main className="flex-1 p-4 md:p-6 space-y-6 overflow-y-auto">
        
        <div className="flex justify-between items-center">
          <h1 className="text-2xl md:text-3xl font-bold">Chats</h1>
          <MainThemeIcon />
        </div>

        {/* Direct Messages */}
        <section className={"rounded-xl shadow p-4 md:p-5 " + CardThemes[theme]}>
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
                    <h3 className="font-semibold pt-1">{msg.name}</h3>
                    <p className="text-sm text-gray-600 dark:text-gray-400">{msg.message}</p>
                  </div>
                </div>
                <span className="text-sm text-gray-500 dark:text-gray-400">{msg.time}</span>
              </div>
            ))}
          </div>
        </section>
        <section className={"rounded-xl shadow p-4 md:p-5 " + CardThemes[theme]}>
          <h2 className="text-lg md:text-xl font-semibold mb-4">Group Messages</h2>
          <div className="space-y-4">
            {groupMessages.map((msg) => (
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
                    <h3 className="font-semibold pt-1">{msg.name}</h3>
                    <p className="text-sm text-gray-600 dark:text-gray-400">{msg.message}</p>
                  </div>
                </div>
                <span className="text-sm text-gray-500 dark:text-gray-400">{msg.time}</span>
              </div>
            ))}
          </div>
        </section>
      </main>
    </div>
  );
}

function NavItem({icon , text}: {icon: React.ReactNode; text: string}) {
  const theme = useRecoilValue(MainThemeAtom);
  const themes = {
    "Bright": "hover:bg-purple-700",
    "Dark": "hover:bg-gray-600",
  }
  return (
    
    <div className={"flex items-center justify-center md:justify-start space-x-0 md:space-x-3 cursor-pointer p-2 md:p-3 rounded-lg " + themes[theme]}>
      {icon}
      <span className="hidden md:inline">{text}</span>
    </div>
  );
}
