import { GroupMessageSelectAtom } from "@/store/atoms/groupMessage_atop";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { useRecoilValue } from "recoil";
import { MessageTypeModalButton } from "../Modals/MessageTypeModal";
import { SendIcon } from "@/icons/send_icon";
import { Input } from "../ui/input";

const CardThemes = {
  Bright: "bg-purple-100",
  Dark: "bg-gray-700",
};

export function GroupMessageSection() {
  const GMSProp = useRecoilValue(GroupMessageSelectAtom);
  const theme = useRecoilValue(MainThemeAtom);

  if (!GMSProp.Valid) {
    return (
      
      <section
        className={`rounded-2xl w-full flex flex-col justify-between gap-5 shadow p-4 md:p-5 ${CardThemes[theme]}`}
      >


      </section>
    );
  }


  return (
    <section
      className={`rounded-2xl w-full flex flex-col justify-between gap-5 shadow p-4 md:p-5 ${CardThemes[theme]}`}>
      
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
