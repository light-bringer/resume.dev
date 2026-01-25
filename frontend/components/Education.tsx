import { Education as EducationType } from '@/types/resume'

interface EducationProps {
  education: EducationType[]
}

export default function Education({ education }: EducationProps) {
  return (
    <section className="card card-hover p-6">
      <h2 className="text-xl font-bold text-slate-900 dark:text-slate-100 mb-4 border-b border-slate-200 dark:border-slate-700 pb-2">
        Education
      </h2>

      <div className="space-y-4">
        {education.map((edu, index) => (
          <div key={index}>
            <h3 className="text-lg font-semibold text-slate-900 dark:text-slate-100">
              {edu.degree}
            </h3>
            <p className="text-blue-600 dark:text-blue-400 font-medium">
              {edu.field}
            </p>
            <div className="flex flex-wrap items-center gap-x-3 gap-y-1 mt-1 text-sm text-slate-600 dark:text-slate-400">
              <span>{edu.institution}</span>
              <span>•</span>
              <span>{edu.startYear} - {edu.endYear}</span>
              <span>•</span>
              <span className="font-medium">{edu.grade}</span>
            </div>
          </div>
        ))}
      </div>
    </section>
  )
}
