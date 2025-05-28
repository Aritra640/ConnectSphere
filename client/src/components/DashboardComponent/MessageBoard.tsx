import { useRecoilValue, useSetRecoilState } from "recoil";
import { MainThemeIcon } from "../global/MainThemeIcon";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { PersonalMessageSearch } from "./PersonalMessageSearchBar";
import { PersonalMessageSection } from "./PersonalMessageSection";
import { PersonalMessageSelectAtom } from "@/store/atoms/personalMessage_atom";

// Tailwind theme classes
const SectionTheme = {
  Bright: "bg-white",
  Dark: "bg-gray-800",
};

const MessageBarThemes = {
  Bright: "bg-purple-100",
  Dark: "bg-gray-700",
};

const MessageBarCompTheme = {
  Bright: "bg-white hover:bg-purple-200",
  Dark: "bg-gray-800 hover:bg-gray-600",
};

const messages = [
  {
    Sender: "Anna",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  },
  {
    Sender: "Penn",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  },
  {
    Sender: "Abc",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  },
  {
    Sender: "Anena",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  },
  {
    Sender: "John",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  },
  {
    Sender: "Jane",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  },
  {
    Sender: "Doe",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  }, {
    Sender: "John",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  },
  {
    Sender: "Jane",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  },
  {
    Sender: "Doe",
    SenderProfile: "/profile.jpg",
    Message: "Hi hello",
    SendAt: "Sunday",
  },

];

export function Messageboard() {
  const theme = useRecoilValue(MainThemeAtom);

  return (
    <main className="flex-1 p-4 md:p-6 space-y-6 overflow-y-auto">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl md:text-3xl font-bold">Personal Messages</h1>
        <MainThemeIcon />
      </div>

      <section
        className={`rounded-xl overflow-hidden flex gap-6 h-[calc(100vh-100px)] shadow p-4 md:p-5 ${SectionTheme[theme]}`}
      >
        <aside
          className={`w-48 h-full md:w-72 flex flex-col gap-4 p-4 md:p-6 rounded-xl overflow-y-auto hide-scrollbar ${MessageBarThemes[theme]}`}
        >
          <PersonalMessageSearch />
          {messages.map((msg, index) => (
            <MessageBarComponent
              key={index}
              Sender={msg.Sender}
              SendAt={msg.SendAt}
              Message={msg.Message}
              SenderProfile={msg.SenderProfile}
            />
          ))}
        </aside>

        <PersonalMessageSection />
      </section>
    </main>
  );
}

interface MessageBarCompProp {
  Sender: string;
  SenderProfile: string;
  Message: string;
  SendAt: string;
}

function MessageBarComponent({
  Sender,
  SenderProfile,
  Message,
  SendAt,
}: MessageBarCompProp) {
  const theme = useRecoilValue(MainThemeAtom);
  const setPersonalMessage = useSetRecoilState(PersonalMessageSelectAtom);

  function SelectPerson() {
    setPersonalMessage({
      Valid: true,
      Username: Sender,
      UserId: 123,
    });
  }

  return (
    <div
      onClick={SelectPerson}
      className={`rounded-xl px-3 py-2 flex items-center gap-3 transition-colors cursor-pointer ${MessageBarCompTheme[theme]}`}
    >
      <Avatar className="w-10 h-10">
        <AvatarImage src={SenderProfile} alt={Sender} />
        <AvatarFallback>
          {Sender.split(" ")
            .map((part) => part[0])
            .join("")
            .toUpperCase()}
        </AvatarFallback>
      </Avatar>

      <div className="flex flex-col">
        <h3 className="font-semibold text-sm">{Sender}</h3>
        <p className="text-xs text-gray-600 dark:text-gray-300 truncate max-w-[120px]">
          {Message}
        </p>
        <span className="text-[10px] text-gray-500 dark:text-gray-400">
          {SendAt}
        </span>
      </div>
    </div>
  );
}

