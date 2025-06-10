import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { useRecoilValue } from "recoil";
import { Button } from "../ui/button";

const Themes = {
  "Bright": "bg-white text-black hover:bg-gray-200",
  "Dark": "bg-gray-800 text-white hover:bg-gray-900",
};


export function NewGroup() {
  const theme = useRecoilValue(MainThemeAtom);

  function NewGroup() {}

  return <Button onClick={NewGroup} className={"rounded-xl px-3 py-6 flex items-center gap-3 transition-colors cursor-pointer " + Themes[theme]}>
    
    NewGroup

  </Button>
}
