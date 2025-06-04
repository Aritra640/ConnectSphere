import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { useRecoilValue } from "recoil";
import { MainThemeIcon } from "../global/MainThemeIcon";

const SectionTheme = {
  Bright: "bg-white",
  Dark: "bg-gray-800",
};

const MessageBarThemes = {
  Bright: "bg-purple-100",
  Dark: "bg-gray-700",
};

export function Groupboard() {
  const theme = useRecoilValue(MainThemeAtom);

  return (
    <main className="flex-1 p-4 md:p-6 space-y-6 overflow-y-auto">
      <div  className="flex justify-between items-center">
        <h1 className="text-2xl md:text-3xl font-bold">Group Messages</h1>
        <MainThemeIcon />
      </div>

      
      <section className={"rounded-xl overflow-hidden flex gap-6 h-[calc(100vh-100px)] shadow p-4 md:p-5 " + SectionTheme[theme]}>

        <aside className={"w-48 h-full md:w-72 flex flex-col gap-4 p-4 md:p-6 rounded-xl overflow-y-auto hide-scrollbar " + MessageBarThemes[theme]}>
          
        </aside>
      </section>

    </main>
  )
}


interface GroupMessageBarComponentProp {
  GroupName: string;
  GroupProfile: string;
  LastMessage: string;
  LastMessageSendAt: string;
}

function GroupMessageBarComponent({
  GroupName,
  GroupProfile,
  LastMessage,
  LastMessageSendAt,
}: GroupMessageBarComponentProp) {
  const theme = useRecoilValue(MainThemeAtom);


}
