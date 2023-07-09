export default function Layout({ children }) {
    return (
        <div className="min-h-screen bg-violet-100 p-16 w-full">
            <div className="bg-slate-100 rounded-3xl grid lg:grid-cols-2 max-w-[72rem] mx-auto overflow-clip">
                <div className="right flex flex-col justify-evenly bg-white">
                    <div className="text-center my-24">
                        {children}
                    </div>
                </div>
                <div class="flex items-center justify-evenly">
                    <img src="\assets\Frame 2 (1).png" class="aspect-auto max-h-64 max-w-64"></img>
                </div>
            </div>

        </div>
    )
}