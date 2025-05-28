import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { useRecoilValue } from "recoil";
import { Avatar, AvatarFallback } from "../ui/avatar";
import { AvatarImage } from "@radix-ui/react-avatar";
import { Input } from "../ui/input";
import { ScrollArea } from "../ui/scroll-area";
import { SendIcon } from "@/icons/send_icon";
import { MessageBox } from "../MessageComponent/MessageBox";
import {
  MessageTypeModal,
  MessageTypeModalButton,
} from "../Modals/MessageTypeModal";
import { PersonalMessageSelectAtom } from "@/store/atoms/personalMessage_atom";

const PersonalChats = [
  { Sender: "Anna", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
  { Sender: "You", Message: "hey there", SendAt: "Sunday" },
];

const CardThemes = {
  Bright: "bg-purple-100",
  Dark: "bg-gray-700",
};

const UserBlockThemes = {
  Bright: "bg-white",
  Dark: "bg-gray-800",
};

const MessageAlign = {
  type1: "flex justify-start",
  type2: "flex justify-end",
};

export function PersonalMessageSection() {
  const PMSProp = useRecoilValue(PersonalMessageSelectAtom);
  const theme = useRecoilValue(MainThemeAtom);

  if (!PMSProp.Valid) {
    return (
      <section
        className={`rounded-2xl w-full flex flex-col justify-between gap-5 shadow p-4 md:p-5 ${CardThemes[theme]}`}
      ></section>
    );
  }

  return (
    <section
      className={`rounded-2xl w-full flex flex-col justify-between gap-5 shadow p-4 md:p-5 ${CardThemes[theme]}`}
    >
      <div
        className={`flex justify-start gap-3 rounded-2xl p-2 ${UserBlockThemes[theme]}`}
      >
        <Avatar>
          <AvatarImage src={PMSProp.Username} alt={PMSProp.Username} />
          <AvatarFallback>
            {PMSProp.Username.split(" ")
              .map((part) => part[0])
              .join("")
              .toUpperCase()}
          </AvatarFallback>
        </Avatar>

        <div className="flex items-center text-xl">{PMSProp.Username}</div>
      </div>

      <ScrollArea className="flex-1 rounded-md border-transparent p-4 overflow-y-auto">
        <div className="flex flex-col gap-2">
          {PersonalChats.map((msg, idx) => (
            <div
              key={idx}
              className={
                msg.Sender === "You" ? MessageAlign.type2 : MessageAlign.type1
              }
            >
              <MessageBox
                Sender={msg.Sender}
                SendAt={msg.SendAt}
                Message={msg.Message}
              />
            </div>
          ))}
        </div>
      </ScrollArea>

      <div>
        <InputMessage />
      </div>
      <MessageTypeModal />
    </section>
  );
}

const InputMessageThemes = {
  Bright: "bg-white hover:bg-gray-300",
  Dark: "bg-gray-800 hover:bg-gray-900",
};

function InputMessage() {
  const theme = useRecoilValue(MainThemeAtom);

  return (
    <div
      className={`flex items-center gap-1 rounded-xl cursor-pointer p-2 ${InputMessageThemes[theme]}`}
    >
      <Input
        className="border-transparent"
        type="text"
        placeholder="Type your message here"
      />
      <MessageTypeModalButton />
      <SendIcon />
    </div>
  );
}
