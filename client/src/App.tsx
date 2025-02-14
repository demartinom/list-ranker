import { useEffect, useState } from "react";
import { getPremades, sendChoice } from "./api/api";

export default function App() {
  const [premadeLists, setPremadeLists] = useState([]);

  useEffect(() => {
    const fetchPremades = async () => {
      const data = await getPremades();
      if (data) setPremadeLists(data.premades);
    };
    fetchPremades();
  }, []);

  const premadeOptions = premadeLists.map((item) => (
    <button key={item} onClick={() => sendChoice(item)}>
      {item}
    </button>
  ));

  return <div>{premadeLists.length > 0 && premadeOptions}</div>;
}
