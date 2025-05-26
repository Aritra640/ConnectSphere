import { MainThemeAtom } from "@/store/atoms/maintheme_atom";
import { useRecoilValue } from "recoil";

const CardThemes = {
  "Bright": "bg-purple-100",
  "Dark": "bg-gray-700",
};

interface PersonalMessageSectionProp {
  Sender: string;
  SenderProfile: string;
}

export function PersonalMessageSection({Sender,SenderProfile}: PersonalMessageSectionProp) {
  const theme = useRecoilValue(MainThemeAtom);

  return <section className={"rounded-2xl shadow p-4 md:p-5 " + CardThemes[theme]}>
    
  </section>
}
