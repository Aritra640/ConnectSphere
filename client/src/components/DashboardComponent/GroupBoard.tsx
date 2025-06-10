import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { useRecoilValue, useSetRecoilState } from "recoil";
import { MainThemeIcon } from "../global/MainThemeIcon";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { GroupMessageSelectAtom } from "@/store/atoms/groupMessage_atop";
import { GroupMessageSection } from "./GroupMessageSection";
import { GroupMessageSearch } from "./GroupSearch";
import { NewGroup } from "./NewGroup";

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

const groupbardummy = [
  {
  GroupName: "Group1",
  GroupProfile: "edqwed",
  LastMessage: "Hibhello there",
  LastMessageSendAt: "Sunday",
  },
  {
  GroupName: "Group1",
  GroupProfile: "edqwed",
  LastMessage: "Hibhello there",
  LastMessageSendAt: "Sunday",
  },
  {
  GroupName: "Group1",
  GroupProfile: "edqwed",
  LastMessage: "Hibhello there",
  LastMessageSendAt: "Sunday",
  },
];

export function Groupboard() {
  const theme = useRecoilValue(MainThemeAtom);

  return (
    <main className="flex-1 p-4 md:p-6 space-y-6 overflow-y-auto">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl md:text-3xl font-bold">Group Messages</h1>
        <MainThemeIcon />
      </div>

      <section
        className={
          "rounded-xl overflow-hidden flex gap-6 h-[calc(100vh-100px)] shadow p-4 md:p-5 " +
          SectionTheme[theme]
        }
      >
        <aside
          className={
            "w-48 h-full md:w-72 flex flex-col gap-4 p-4 md:p-6 rounded-xl overflow-y-auto hide-scrollbar " +
            MessageBarThemes[theme]
          }
        >
          <GroupMessageSearch />
          <NewGroup />
          {groupbardummy.map((msg, index) => (
            <GroupMessageBarComponent 
              key={index}
              GroupName={msg.GroupName}
              GroupProfile={msg.GroupProfile}
              LastMessage={msg.LastMessage}
              LastMessageSendAt={msg.LastMessageSendAt} />
          ))}
        </aside>

        <GroupMessageSection />
      </section>
    </main>
  );
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
  const setGroupMessage = useSetRecoilState(GroupMessageSelectAtom);

  function SelectGroup() {
    setGroupMessage({
      Valid: true,
      GroupName: GroupName,
      GroupAvatar: GroupProfile,
      GroupId: "123",
    });
  }

  return (
    <div
      onClick={SelectGroup}
      className={
        "rounded-xl px-3 py-2 flex items-center gap-3 transtition-colors cursor-pointer " +
        MessageBarCompTheme[theme]
      }
    >
      <Avatar className="w-10 h-10">
        <AvatarImage src={GroupProfile} alt={GroupName} />
        <AvatarFallback>
          {GroupName.split(" ")
            .map((part) => part[0])
.join("")
            .toUpperCase()}
        </AvatarFallback>
      </Avatar>

      <div className="flex flex-col">
        <h3 className="font-semibold text-sm">{GroupName}</h3>
        <p className="text-xs text-gray-600 dark:text-gray-300 truncate max-w-[120px]">
          {LastMessage}
        </p>

        <span className="text-[10px] text-gray-500 dark:text-gray-400">
          {LastMessageSendAt}
        </span>
      </div>
    </div>
  );
}
