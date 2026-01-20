import { Experience as ExperienceType } from '@/types/resume'

interface ExperienceProps {
  experiences: ExperienceType[]
}

function formatDate(dateString: string | null | undefined): string {
  if (!dateString) return 'Present'
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { month: 'short', year: 'numeric' })
}

export default function Experience({ experiences }: ExperienceProps) {
  return (
    <section className="card card-hover p-6">
      <h2 className="text-xl font-bold text-slate-900 dark:text-slate-100 mb-4 border-b border-slate-200 dark:border-slate-700 pb-2">
        Experience
      </h2>

      <div className="space-y-6">
        {experiences.map((exp, index) => (
          <div key={index} className="relative">
            {/* Job Title & Company */}
            <div className="mb-2">
              <h3 className="text-lg font-semibold text-slate-900 dark:text-slate-100">
                {exp.title}
              </h3>
              <div className="flex flex-wrap items-center gap-x-3 gap-y-1 text-sm text-slate-600 dark:text-slate-400">
                <span className="font-medium text-blue-600 dark:text-blue-400">
                  {exp.company}
                </span>
                <span>•</span>
                <span>{exp.location}</span>
                <span>•</span>
                <span>
                  {formatDate(exp.startDate)} - {formatDate(exp.endDate)}
                </span>
              </div>
              {exp.team && (
                <p className="text-sm text-slate-500 dark:text-slate-500 mt-1">
                  {exp.team}
                </p>
              )}
            </div>

            {/* Tech Stack */}
            <div className="flex flex-wrap gap-2 mb-3">
              {exp.techStack.map((tech, techIndex) => (
                <span
                  key={techIndex}
                  className="px-2 py-1 text-xs font-medium bg-slate-100 dark:bg-slate-700 text-slate-700 dark:text-slate-300 rounded"
                >
                  {tech}
                </span>
              ))}
            </div>

            {/* Achievements */}
            <ul className="space-y-1.5 text-sm text-slate-700 dark:text-slate-300">
              {exp.achievements.map((achievement, achIndex) => (
                <li key={achIndex} className="flex items-start gap-2">
                  <span className="text-blue-600 dark:text-blue-400 mt-1.5 flex-shrink-0">•</span>
                  <span className="leading-relaxed">{achievement}</span>
                </li>
              ))}
            </ul>

            {/* Divider between experiences */}
            {index < experiences.length - 1 && (
              <div className="mt-6 border-b border-slate-200 dark:border-slate-700" />
            )}
          </div>
        ))}
      </div>
    </section>
  )
}
