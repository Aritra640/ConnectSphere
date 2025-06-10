import { useRecoilState } from "recoil";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { Switch } from "@/components/ui/switch";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
  Moon,
  Sun,
  Bell,
  Shield,
  LogOut,
  Lock,
  UserCircle,
} from "lucide-react";
import { useState } from "react";

const Themes = {
  Bright: "bg-white",
  Dark: "bg-slate-950",
};

export function SettingsBoard() {
  const [theme, setTheme] = useRecoilState(MainThemeAtom);
  const isDark = theme === "Dark";

  const [notification, setNotification] = useState<true | false>(true);

  const toggleNotification = () => {
    setNotification((p) => !p);
  };

  const toggleTheme = () => {
    setTheme(isDark ? "Bright" : "Dark");
  };

  return (
    <div
      className={
        "w-full px-4 py-8 flex flex-col items-center gap-6 bg-muted min-h-screen " +
        Themes[theme]
      }
    >
      <h1 className="text-3xl font-bold tracking-tight">Settings</h1>

      {/* Account Card */}
      <Card className="w-full max-w-2xl rounded-2xl shadow-md">
        <CardHeader className="flex flex-row items-center gap-3">
          <div className="flex items-center gap-3">
          <UserCircle className="text-primary" />
          <CardTitle className="text-xl">Account</CardTitle>
          </div>
          <Button variant="outline" className="text-sm cursor-pointer px-3 py-1">
            Edit
          </Button>
        </CardHeader>
        <CardContent className="flex items-center gap-4 px-6 pb-6">
          <Avatar className="w-14 h-14">
            <AvatarImage src="/user.png" alt="User" />
            <AvatarFallback>AC</AvatarFallback>
          </Avatar>
          <div>
            <p className="font-medium text-lg">Aritra Chatterjee</p>
            <p className="text-muted-foreground text-sm">aritra@example.com</p>
          </div>
        </CardContent>
      </Card>

      {/* Theme Card */}
      <Card className="w-full max-w-2xl rounded-2xl shadow-md">
        <CardHeader className="flex flex-row items-center gap-3">
          {isDark ? (
            <Moon className="text-primary" />
          ) : (
            <Sun className="text-primary" />
          )}
          <CardTitle className="text-xl">Appearance</CardTitle>
        </CardHeader>
        <CardContent className="flex items-center justify-between px-6 pb-6">
          <span className="text-base">Dark Mode</span>
          <Switch
            className="cursor-pointer"
            checked={isDark}
            onCheckedChange={toggleTheme}
          />
        </CardContent>
      </Card>

      {/* Notifications Card */}
      <Card className="w-full max-w-2xl rounded-2xl shadow-md">
        <CardHeader className="flex flex-row items-center gap-3">
          <Bell className="text-primary" />
          <CardTitle className="text-xl">Notifications</CardTitle>
        </CardHeader>
        <CardContent className="flex items-center justify-between px-6 pb-6">
          <span className="text-base">Enable Notifications</span>
          <Switch
            className="cursor-pointer"
            checked={notification}
            onCheckedChange={toggleNotification}
          />
        </CardContent>
      </Card>

      {/* Security Card */}
      <Card className="w-full max-w-2xl rounded-2xl shadow-md">
        <CardHeader className="flex flex-row items-center gap-3">
          <Shield className="text-primary" />
          <CardTitle className="text-xl">Security</CardTitle>
        </CardHeader>
        <CardContent className="flex flex-col gap-4 px-6 pb-6">
          <Button
            variant="outline"
            className="w-full flex items-center gap-2 cursor-pointer hover:bg-muted"
          >
            <Lock className="w-4 h-4" /> Change Password
          </Button>
          <Button
            variant="destructive"
            className="w-full cursor-pointer flex items-center gap-2 hover:brightness-110"
          >
            <LogOut className="w-4 h-4" /> Logout
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
