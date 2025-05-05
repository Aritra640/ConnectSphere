import { useRecoilValue, useSetRecoilState } from "recoil"
import { TestAtom } from "../store/atoms/test_atom"


export function TestComponent() {

  const BulbVal = useRecoilValue(TestAtom);
  const setBulb = useSetRecoilState(TestAtom);

  function ToggleBulb() {

    setBulb(p => !p);
  };

  return <div> 
    <div>

    {
      BulbVal ? <span>On</span> : <span>Off</span>
    }
  </div>

  <button onClick={ToggleBulb}>Toggle</button>

  </div>
}
