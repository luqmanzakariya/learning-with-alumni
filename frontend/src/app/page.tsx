export default async function Home() {
  try {
    const res = await fetch("http://localhost:8080/users");
    const data = await res.json();

    return (
      <div className="font-sans grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 bg-gray-50">
        <div>
          <div className="pt-[121px] text-2xl font-bold text-blue-700 mb-6 text-center">
            Student Data
          </div>
          <div className="flex w-full gap-4">
            {data.map((val: any, i: number) => (
              <div
                key={`${i}-${val.id}`}
                className="border border-gray-300 rounded-lg shadow-md p-4 bg-white hover:shadow-lg transition-shadow duration-200 w-64 flex flex-col gap-2"
              >
                <div className="text-lg font-semibold text-gray-800">
                  Nama: {val.name}
                </div>
                <div className="text-sm text-gray-600">
                  Email: {val.email}
                </div>
                <div className="text-sm text-gray-600">
                  Alamat: {val.address}
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    );
  } catch (error) {
    return (
      <div className="font-sans grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20">
        <div className="flex w-full gap-4">Under maintenance</div>
      </div>
    );
  }
}
