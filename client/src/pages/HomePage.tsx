import { ThemeAtom } from "@/store/atoms/theme_atom"
import { useRecoilValue } from "recoil"

const Themes = {
  "Dark": "bg-slate-950 text-white",
  "Bright": "bg-white text-black"
}

export function HomePage() {

  const themeVal = useRecoilValue(ThemeAtom);

  return <div className={Themes[themeVal] + " h-screen w-screen"}>
    
    
  </div> 
}
