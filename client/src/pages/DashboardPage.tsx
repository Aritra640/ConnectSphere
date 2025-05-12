import { MainThemeAtom } from "@/store/atoms/maintheme_atom"
import { useRecoilState, useRecoilValue } from "recoil"

const MainTheme = {

  "Bright": "bg-slate-100",
  "Dark": "bg-slate-950",
}

export function DashboardPage() {

  const mainTheme = useRecoilValue(MainThemeAtom);

  return <div className={MainTheme[mainTheme] + " h-screen v-screen "}>

        
  </div>
}
