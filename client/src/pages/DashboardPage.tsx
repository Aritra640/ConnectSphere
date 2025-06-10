import { Groupboard } from "@/components/DashboardComponent/GroupBoard";
import { Homeboard } from "@/components/DashboardComponent/HomeBoard";
import { Messageboard } from "@/components/DashboardComponent/MessageBoard";
import { SettingsBoard } from "@/components/DashboardComponent/SettingBoard";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { DashBoardAtom } from "@/store/atoms/dashboard_atom";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { Home, LogOut, MessageCircle, Settings, Users } from "lucide-react";
import { useRecoilValue, useSetRecoilState } from "recoil";

const DashboardTheme = {
  "Bright": "bg-gray-200 text-black",
  "Dark": "bg-gray-950 text-white",
};

const SidebarTheme = {
  "Bright": "bg-purple-800",
  "Dark": "bg-gray-800",
};



export function DashboardPage() {
  const defaultColorTheme = "flex h-screen transition-colors duration-300";
  const theme = useRecoilValue(MainThemeAtom);
  const dashboardState = useRecoilValue(DashBoardAtom);
  const setDashboardState = useSetRecoilState(DashBoardAtom);

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
              theme === "Dark" ? "bg-gray-700 hover:bg-gray-600" : "bg-purple-700 hover:bg-purple-600"
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
            <NavItem icon={<Home />} text="Home" onClick={()=>{setDashboardState("Home")}}/>
            <NavItem icon={<MessageCircle />} text="Messages" onClick={()=>{setDashboardState("Message")}}/>
            <NavItem icon={<Users />} text="Groups" onClick={()=>{setDashboardState("Group")}}/>
            <NavItem icon={<Settings />} text="Settings" onClick={()=>{setDashboardState("Setting")}}/>
            <NavItem icon={<LogOut />} text="Logout" onClick={()=>{setDashboardState("Setting")}}/>
          </nav>
        </div>
      </aside>

      <>
        {dashboardState == "Home" && <Homeboard />}
        {dashboardState == "Message" && <Messageboard />}
        {dashboardState == "Group" && <Groupboard />}
        {dashboardState == "Setting" && <SettingsBoard />}
      </>

    </div>
  );
}

function NavItem({icon , text, onClick}: {icon: React.ReactNode; text: string; onClick: ()=>void}) {
  const theme = useRecoilValue(MainThemeAtom);
  const themes = {
    "Bright": "hover:bg-purple-700",
    "Dark": "hover:bg-gray-600",
  }
  return (
    
    <div onClick={onClick} className={"flex items-center justify-center md:justify-start space-x-0 md:space-x-3 cursor-pointer p-2 md:p-3 rounded-lg " + themes[theme]}>
      {icon}
      <span className="hidden md:inline">{text}</span>
    </div>
  );
}

