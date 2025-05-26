import { useRecoilValue } from "recoil";
import { MainThemeIcon } from "../global/MainThemeIcon";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";

const HoverMessages = {
  "Bright": "hover:bg-purple-300",
  "Dark": "hover:bg-gray-600",
};

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

const CardThemes = {
  "Bright": "bg-white",
  "Dark": "bg-gray-800",
};

export function Homeboard() {
  
  const theme = useRecoilValue(MainThemeAtom);

  return (
    <main className="flex-1 p-4 md:p-6 space-y-6 overflow-y-auto">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl md:text-3xl font-bold">Chats</h1>
        <MainThemeIcon />
      </div>

      {/* Direct Messages */}
      <section className={"rounded-xl shadow p-4 md:p-5 " + CardThemes[theme]}>
        <h2 className="text-lg md:text-xl font-semibold mb-4">
          Direct Messages
        </h2>
        <div className="space-y-4">
          {messages.map((msg) => (
            <div
              key={msg.name}
              className={
                "flex justify-between items-start rounded-xl cursor-pointer p-2 " +
                HoverMessages[theme]
              }
            >
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
                  <p className="text-sm text-gray-500 dark:text-gray-300">
                    {msg.message}
                  </p>
                </div>
              </div>
              <span className="text-sm text-gray-500 dark:text-gray-400">
                {msg.time}
              </span>
            </div>
          ))}
        </div>
      </section>
      <section className={"rounded-xl shadow p-4 md:p-5 " + CardThemes[theme]}>
        <h2 className="text-lg md:text-xl font-semibold mb-4">
          Group Messages
        </h2>
        <div className="space-y-4">
          {groupMessages.map((msg) => (
            <div
              key={msg.name}
              className={
                "flex justify-between items-start rounded-xl cursor-pointer p-2 " +
                HoverMessages[theme]
              }
            >
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
                  <p className="text-sm text-gray-500 dark:text-gray-300">
                    {msg.sender + ": " + msg.message}
                  </p>
                </div>
              </div>
              <span className="text-sm text-gray-500 dark:text-gray-400">
                {msg.time}
              </span>
            </div>
          ))}
        </div>
      </section>
    </main>
  );
}
