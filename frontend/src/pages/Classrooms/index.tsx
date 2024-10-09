import TopNav from "@/components/Layout/TopNav";

const Classrooms: React.FC = () => {
    
    return(
        <div className="Classrooms">
            <div>
                <TopNav/>
            </div>
            <div className="Classrooms__bottom">
                <div className="Classrooms__content">
                    <div className="Classrooms__title">Classrooms</div>
                </div>
            </div>
        </div>
    )
} 

export default Classrooms;