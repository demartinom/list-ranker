import { useEffect, useState } from "react";
import { getPremades } from "./api/api";

export default function App() {
  const [premadeLists, setPremadeLists] = useState([]);

  useEffect(() => {
    const fetchPremades = async () => {
      const data = await getPremades();
      if (data) setPremadeLists(data);
    };
    fetchPremades();
  }, []);
  console.log(premadeLists);
  return <div>App</div>;
}
