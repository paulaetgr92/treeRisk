import { useEffect, useState } from "react";
import TreeForm from "./components/TreeForm";
import TreeList from "./components/TreeList";
import Navbar from "./components/Navbar";
import { getTrees } from "./service/api";
import "./styles.css";

function App() {
    const [trees, setTrees] = useState([]);

    const fetchTrees = async (newTree) => {
        if (newTree) {
            setTrees((prev) => [newTree, ...prev]);
            return;
        }

        const data = await getTrees();
        setTrees(Array.isArray(data) ? data : []);
    };

    useEffect(() => {
        fetchTrees();
    }, []);

return (
        <>
            {/* 🔥 NAVBAR AQUI */}
            <Navbar />

            <div className="container">
                <h1>Tree Risk Monitor</h1>

                <TreeForm onCreated={fetchTrees} />
                <TreeList trees={trees} />
            </div>
        </>
    );
}

export default App;