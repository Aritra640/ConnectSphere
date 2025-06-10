import { SearchIcon } from "@/icons/search_icon";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom"; import { useRecoilValue } from "recoil";
import { Input } from "../ui/input";

const Themes = {
  "Bright": "bg-white text-black hover:bg-gray-200",
  "Dark": "bg-gray-800 text-white hover:bg-gray-900",
};

export function GroupMessageSearch() {
  const theme = useRecoilValue(MainThemeAtom);

  return <div className={"rounded-xl px-3 py-2 flex justify-between gap-3 transition-colors  cursor-pointer " + Themes[theme]}>

    <InputComponent />
    <SearchGroup />

  </div>
}

function SearchGroup() {

  return <div className="cursor-pointer flex items-center border-transparent"><SearchIcon /></div>
}


function InputComponent() {

  return <Input className="border-transparent" type="email" placeholder="search by email" />
}
