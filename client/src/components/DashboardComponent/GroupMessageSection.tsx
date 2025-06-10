import { useRecoilValue } from "recoil";
import { GroupMessageSelectAtom } from "@/store/atoms/groupMessage_atop";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import {
  MessageTypeModal,
  MessageTypeModalButton,
} from "../Modals/MessageTypeModal";
import { SendIcon } from "@/icons/send_icon";
import { Input } from "../ui/input";
import { ScrollArea } from "../ui/scroll-area";
import { Avatar, AvatarFallback } from "../ui/avatar";
import { AvatarImage } from "@radix-ui/react-avatar";
import { MessageBox } from "../MessageComponent/MessageBox";

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

const InputMessageThemes = {
  Bright: "bg-white hover:bg-gray-300",
  Dark: "bg-gray-800 hover:bg-gray-900",
};

// Dummy messages for demo
const GroupChats = [
  { Sender: "Anna", Message: "Hey team", SendAt: "Monday" },
  { Sender: "You", Message: "Hello everyone", SendAt: "Monday" },
  { Sender: "John", Message: "What’s the update?", SendAt: "Monday" },
];

export function GroupMessageSection() {
  const GMSProp = useRecoilValue(GroupMessageSelectAtom);
  const theme = useRecoilValue(MainThemeAtom);

  if (!GMSProp.Valid) {
    return (
      <section
        className={`rounded-2xl w-full flex flex-col justify-between gap-5 shadow p-4 md:p-5 ${CardThemes[theme]}`}
      />
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
          <AvatarImage
            src={GMSProp.GroupAvatar || ""}
            alt={GMSProp.GroupName}
          />
          <AvatarFallback>
            {GMSProp.GroupName?.split(" ")
              .map((part) => part[0])
              .join("")
              .toUpperCase()}
          </AvatarFallback>
        </Avatar>
        <div className="flex items-center text-xl font-medium truncate max-w-[70%] sm:max-w-[85%]">{GMSProp.GroupName}</div>
      </div>

      <ScrollArea className="flex-1 rounded-md border-transparent p-4 overflow-y-auto">
        <div className="flex flex-col gap-2">
          {GroupChats.map((msg, idx) => (
            <div
              key={idx}
              className={
                msg.Sender === "You" ? MessageAlign.type2 : MessageAlign.type1
              }
            >
              <MessageBox
                Sender={msg.Sender}
                Message={msg.Message}
                SendAt={msg.SendAt}
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
