import { useRecoilValue } from "recoil";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { MoreVertical } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

interface MessageBoxProp {
  Sender: string;
  Message: string;
  SendAt: string;
}

const MessageBoxThemes = {
  Bright: "bg-purple-200",
  Dark: "bg-gray-800",
};

const YouThemes = {
  Bright: "bg-white text-black",
  Dark: "bg-gray-400 text-black",
};

export function MessageBox({ Sender, Message, SendAt }: MessageBoxProp) {
  const theme = useRecoilValue(MainThemeAtom);
  const isYou = Sender === "You";

  return (
    <div
      className={
        "relative w-fit max-w-sm sm:max-w-md md:max-w-lg xl:max-w-2xl px-4 py-2 rounded-xl shadow-md " +
        (isYou ? YouThemes[theme] : MessageBoxThemes[theme])
      }
    >
      {/* Header: Sender, SendAt, and 3-dot menu in one line */}
      <div className="flex items-center justify-between gap-2">
        <div className="flex items-center gap-2 text-sm font-semibold">
          <span>{Sender}</span>
          <span className="text-xs text-gray-500">{SendAt}</span>
        </div>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <button className="p-1 rounded-md hover:bg-gray-300 cursor-pointer dark:hover:bg-gray-600">
              <MoreVertical className="w-4 h-4" />
            </button>
          </DropdownMenuTrigger>
          <DropdownMenuContent side="bottom" align="end" className="w-32">
            <DropdownMenuItem onClick={() => alert("Edit")}>Edit</DropdownMenuItem>
            <DropdownMenuItem onClick={() => alert("Delete")}>Delete</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>

      {/* Message */}
      <p className="text-sm mt-1">{Message}</p>
    </div>
  );
}

