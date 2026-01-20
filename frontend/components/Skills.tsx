import { TechnicalSkills } from '@/types/resume'

interface SkillsProps {
  skills: TechnicalSkills
}

interface SkillCategoryProps {
  title: string
  items: string[]
}

function SkillCategory({ title, items }: SkillCategoryProps) {
  return (
    <div>
      <h3 className="text-sm font-semibold text-slate-900 dark:text-slate-100 mb-2">
        {title}
      </h3>
      <div className="flex flex-wrap gap-2">
        {items.map((item, index) => (
          <span
            key={index}
            className="px-2 py-1 text-xs font-medium bg-slate-100 dark:bg-slate-700 text-slate-700 dark:text-slate-300 rounded"
          >
            {item}
          </span>
        ))}
      </div>
    </div>
  )
}

export default function Skills({ skills }: SkillsProps) {
  return (
    <section className="card card-hover p-6">
      <h2 className="text-xl font-bold text-slate-900 dark:text-slate-100 mb-4 border-b border-slate-200 dark:border-slate-700 pb-2">
        Technical Skills
      </h2>

      <div className="grid gap-4 sm:grid-cols-2">
        <SkillCategory title="Languages" items={skills.languages} />
        <SkillCategory title="Databases & Messaging" items={skills.databases} />
        <SkillCategory title="Cloud & Infrastructure" items={skills.cloudInfra} />
        <SkillCategory title="Specializations" items={skills.specializations} />
        <div className="sm:col-span-2">
          <SkillCategory title="Version Control" items={skills.versionControl} />
        </div>
      </div>
    </section>
  )
}
