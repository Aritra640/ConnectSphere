import { MainThemeAtom } from "@/store/atoms/maintheme_atom"
import { useRecoilValue } from "recoil"
import { Input } from "../ui/input";
import { SearchIcon } from "@/icons/search_icon";

const Themes = {
  "Bright": "bg-white text-black hover:bg-gray-200",
  "Dark": "bg-gray-800 text-white hover:bg-gray-900",
};

export function PersonalMessageSearch() {
  const theme = useRecoilValue(MainThemeAtom);

  return <div className={"rounded-xl px-3 py-2 flex justify-between gap-3 transition-colors cursor-pointer " + Themes[theme]}>
    
    <InputComponent />
    <SearchPersonalUser />
  </div>
}

function SearchPersonalUser() {

  return <div className="cursor-pointer flex items-center border-transparent"><SearchIcon /></div>
}


function InputComponent() {

  return <Input className="border-transparent" type="email" placeholder="search by email" />
}
