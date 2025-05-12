import { Button } from "../ui/button"

interface ButtonComponentProp {
  Type: "first"|"second"|"third",
  Content: string, 
  OnClick: () => void,
}

const Themes = {
  "first": "bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-xl font-medium transition cursor-pointer",
  "second": "border border-indigo-600 text-indigo-400 hover:text-white hover:border-white px-4 py-2 rounded-xl font-medium transition cursor-pointer",
  "third" : "border border-slate-700 hover:bg-slate-800 text-slate-300 hover:text-white py-2 rounded-xl transition font-medium cursor-pointer"
}

export function ButtonComponent({Type , Content , OnClick}: ButtonComponentProp) {

  return <Button onClick={OnClick} className={Themes[Type] + " text-center"}>{Content}</Button>
}
