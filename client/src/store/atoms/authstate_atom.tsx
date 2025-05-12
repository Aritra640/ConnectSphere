import { atom } from "recoil";

export const AuthStateAtom = atom<"Signin"|"Signup">({
  default: "Signin",
  key: "AuthStateAtom",
});
