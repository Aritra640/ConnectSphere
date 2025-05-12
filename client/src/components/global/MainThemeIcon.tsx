import { BrightThemeIcon } from "@/icons/brighttheme_icon";
import { DarkThemeIcon } from "@/icons/darktheme_icon";
import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { useRecoilState } from "recoil";

const Theme = {
  Bright: "text-black",
  Dark: "text-white",
};

export function MainThemeIcon() {
  const [mainTheme, setMainTheme] = useRecoilState(MainThemeAtom);

  function ToggleMainTheme() {
    if (mainTheme == "Bright") setMainTheme("Dark");
    else setMainTheme("Bright");
  }

  return (
    <button
      onClick={ToggleMainTheme}
      className={Theme[mainTheme] + " cursor-pointer"}
    >
      {mainTheme == "Bright" ? <BrightThemeIcon /> : <DarkThemeIcon />}
    </button>
  );
}
