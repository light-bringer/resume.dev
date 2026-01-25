interface SummaryProps {
  summary: string
}

export default function Summary({ summary }: SummaryProps) {
  return (
    <section className="card card-hover p-6">
      <h2 className="text-xl font-bold text-slate-900 dark:text-slate-100 mb-3 border-b border-slate-200 dark:border-slate-700 pb-2">
        About
      </h2>
      <p className="text-slate-700 dark:text-slate-300 leading-relaxed">
        {summary}
      </p>
    </section>
  )
}
