interface InterestsProps {
  interests: string[]
}

export default function Interests({ interests }: InterestsProps) {
  return (
    <section className="card card-hover p-6">
      <h2 className="text-xl font-bold text-slate-900 dark:text-slate-100 mb-4 border-b border-slate-200 dark:border-slate-700 pb-2">
        Interests
      </h2>

      <ul className="grid gap-2 sm:grid-cols-2">
        {interests.map((interest, index) => (
          <li
            key={index}
            className="flex items-start gap-2 text-sm text-slate-700 dark:text-slate-300"
          >
            <span className="text-blue-600 dark:text-blue-400 mt-0.5 flex-shrink-0">•</span>
            <span>{interest}</span>
          </li>
        ))}
      </ul>
    </section>
  )
}
